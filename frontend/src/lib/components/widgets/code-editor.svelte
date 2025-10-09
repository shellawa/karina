<script lang="ts">
  import loader from "@monaco-editor/loader"
  import type * as Monaco from "monaco-editor/esm/vs/editor/editor.api"
  import { onDestroy, onMount } from "svelte"

  let editor: Monaco.editor.IStandaloneCodeEditor
  let monaco: typeof Monaco
  let editorContainer: HTMLElement

  interface Props {
    value: string
    language?: string
    theme?: string
    readonly?: boolean
  }

  let { value = $bindable(), language, theme, readonly = false }: Props = $props()

  onMount(async () => {
    const monacoEditor = await import("monaco-editor")
    loader.config({ monaco: monacoEditor.default })
    monaco = await loader.init()

    editor = monaco.editor.create(editorContainer, {
      value,
      language,
      theme,
      fontSize: 16,
      automaticLayout: true,
      minimap: { enabled: false },
      overviewRulerLanes: 0,
      overviewRulerBorder: false,
      scrollbar: { vertical: "hidden", horizontal: "hidden" },
      readOnly: readonly
    })

    editor.onDidChangeModelContent(() => {
      value = editor.getValue()
    })
  })

  $effect(() => {
    if (editor && editor.getValue() !== value && !editor.hasTextFocus()) {
      editor.setValue(value)
    }
  })
</script>

<div class="container h-[400px] rounded-lg border px-1.5" bind:this={editorContainer}></div>
