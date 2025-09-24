import fs from "fs"
import ky from "ky"
import path from "path"
import { fileURLToPath } from "url"
import yauzl from "yauzl"

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

const python = {
  name: "python-3.13.7",
  url: "https://www.python.org/ftp/python/3.13.7/python-3.13.7-embed-amd64.zip",
  path: "python-3.13.7-embed-amd64"
}
const runtimesPath = path.join("..", "build", "bin", "assets", "runtimes")
const rootPath = path.join("scripts")
const tempPath = path.join(rootPath, "temp")
if (!fs.existsSync(tempPath)) fs.mkdirSync(tempPath, { recursive: true })
if (!fs.existsSync(runtimesPath)) fs.mkdirSync(runtimesPath, { recursive: true })

export async function extractZip(zipPath, outDir) {
  return new Promise((resolve, reject) => {
    yauzl.open(zipPath, { lazyEntries: true }, (err, zipfile) => {
      if (err) return reject(err)

      zipfile.readEntry()
      zipfile.on("entry", (entry) => {
        const destPath = path.join(outDir, entry.fileName)

        if (/\/$/.test(entry.fileName)) {
          fs.mkdir(destPath, { recursive: true }, (err) => {
            if (err) return reject(err)
            zipfile.readEntry()
          })
        } else {
          fs.mkdir(path.dirname(destPath), { recursive: true }, (err) => {
            if (err) return reject(err)

            zipfile.openReadStream(entry, (err, readStream) => {
              if (err) return reject(err)

              const writeStream = fs.createWriteStream(destPath)
              readStream.pipe(writeStream)

              writeStream.on("close", () => {
                zipfile.readEntry()
              })
            })
          })
        }
      })

      zipfile.on("end", resolve)
      zipfile.on("error", reject)
    })
  })
}

const setupPython = async () => {
  const res = await ky.get(python.url)
  const pythonPath = path.join(path.join(tempPath), python.path)
  fs.writeFileSync(pythonPath + ".zip", await res.bytes())
  await extractZip(pythonPath + ".zip", path.join(runtimesPath, "python"))
}

await setupPython()
fs.rmSync(tempPath, { recursive: true, force: true })
