<script lang="ts">
  import { onMount } from "svelte"
  import { FileText, FlaskConical, Trash2 } from "@lucide/svelte"
  import { EventsOn } from "$lib/wailsjs/runtime/runtime"
  import { SelectFile } from "$lib/wailsjs/go/helpers/Service"
  import {
    GetProblems,
    GetTestCases,
    WriteGeneratorScript,
    GetGeneratorScript,
    WriteSolution
  } from "$lib/wailsjs/go/models/Service"
  import { type models } from "$lib/wailsjs/go/models"
  import { Button } from "$lib/components/ui/button"
  import * as Card from "$lib/components/ui/card"
  import { Input } from "$lib/components/ui/input"
  import { Label } from "$lib/components/ui/label"
  import * as Select from "$lib/components/ui/select"
  import CodeEditor from "$lib/components/widgets/code-editor.svelte"
  import AddManualTest from "./addManualTest.svelte"
  import { GenerateTests } from "$lib/wailsjs/go/languages/Service"
  import { DeleteTest } from "$lib/wailsjs/go/models/Service"

  let filePath = $state("")
  async function selectSolutionFile() {
    const path = await SelectFile()
    filePath = path
    await WriteSolution(selectedProblemId, path)
  }

  let testsNum = $state(1)

  let problems = $state<models.Problem[]>([])
  let selectedProblemId = $state("")
  let selectedProblem = $derived(problems.find((x) => x.id === selectedProblemId))

  let dummy = $state(0)
  let testCases = $derived.by(async () => {
    const aaa = dummy
    return await GetTestCases(selectedProblemId)
  })

  let generatorScript = $state("")

  $effect(() => {
    if (selectedProblemId) {
      WriteGeneratorScript(selectedProblemId, generatorScript)
    }
  })

  onMount(async () => {
    problems = await GetProblems()
    if (problems.length > 0) {
      selectedProblemId = problems[0].id
      generatorScript = await GetGeneratorScript(selectedProblemId)
    }
  })

  EventsOn("problem:change", async (id) => {
    problems = await GetProblems()
    selectedProblemId = id
    generatorScript = await GetGeneratorScript(selectedProblemId)
  })

  EventsOn("generate:test_generated", () => {
    dummy++
  })
</script>

<div class="grid grid-cols-2 gap-4">
  <!-- test generating card -->
  <div>
    <Card.Root>
      <Card.Header>
        <Card.Title class="flex items-center gap-x-4">
          <FlaskConical class="h-5 w-5" />
          <div>Test Generator</div>
        </Card.Title>
      </Card.Header>
      <Card.Content>
        <div class="space-y-3">
          <div>
            <Label class="mb-2">Select problem</Label>
            <Select.Root
              type="single"
              bind:value={selectedProblemId}
              onValueChange={async (id) => (generatorScript = await GetGeneratorScript(id))}
            >
              <Select.Trigger>
                {selectedProblem?.label}
              </Select.Trigger>

              <Select.Content>
                {#each problems as problem}
                  <Select.Item value={problem.id}>{problem.label}</Select.Item>
                {/each}
              </Select.Content>
            </Select.Root>
          </div>

          <Label class="mb-2">Generator script</Label>

          {#key selectedProblemId}
            <CodeEditor language="python" theme="vs" bind:value={generatorScript} />
          {/key}

          <div class="grid grid-cols-2">
            <div>
              <Label class="mb-2">Solution (for generating output)</Label>
              <Button variant={filePath ? "default" : "secondary"} onclick={selectSolutionFile}>
                {filePath ? filePath.split("\\").at(-1)?.split("/").at(-1) : "Select file"}
              </Button>
            </div>
            <div>
              <Label class="mb-2">Number of test(s) to generate</Label>
              <Input type="number" bind:value={testsNum} />
            </div>
          </div>
        </div>
      </Card.Content>
      <Card.Footer>
        <Button
          class="ml-auto"
          onclick={async () => await GenerateTests(selectedProblemId, testsNum)}
        >
          Generate
        </Button>
      </Card.Footer>
    </Card.Root>
  </div>

  <!-- generated tests -->
  <div>
    <Card.Root>
      <Card.Header>
        <Card.Title class="flex items-center justify-between">
          <div class="flex items-center gap-x-4">
            <FileText class="h-5 w-5" />
            <div>Current Tests</div>
          </div>
          <AddManualTest />
        </Card.Title>
      </Card.Header>
      <Card.Content>
        {#await testCases then testCases}
          <div class="space-y-3">
            {#each testCases as testCase}
              <div class="rounded-lg border p-3 transition-colors hover:bg-gray-50">
                <div class="mb-2 flex items-center justify-between">
                  <span class="font-medium">{testCase.Id}</span>
                  <Button
                    variant="ghost"
                    size="sm"
                    onclick={() => DeleteTest(selectedProblemId, testCase.Id)}
                  >
                    <Trash2 class="h-4 w-4" />
                  </Button>
                </div>
                <div class="rounded bg-gray-50 p-2 font-mono text-sm">
                  <div class="text-gray-600">Input:</div>
                  <div class="whitespace-pre">{testCase.Input}</div>
                  <div class="mt-2 text-gray-600">Expected Output:</div>
                  <div class="whitespace-pre">{testCase.Output}</div>
                </div>
              </div>
            {/each}
          </div>
        {/await}
      </Card.Content>
    </Card.Root>
  </div>
</div>
