<script lang="ts">
  import { buttonVariants } from "$lib/components/ui/button"
  import * as Dialog from "$lib/components/ui/dialog"
  import { FileText, List } from "@lucide/svelte"
  import CodeEditor from "$lib/components/widgets/code-editor.svelte"

  let { content }: { content: string } = $props()
</script>

<Dialog.Root>
  <Dialog.Trigger class={buttonVariants({ variant: "outline", size: "sm" })}>
    <FileText /> View Code
  </Dialog.Trigger>
  <Dialog.Content class="lg:max-w-4xl">
    <Dialog.Header>
      <Dialog.Title>Code</Dialog.Title>
    </Dialog.Header>
    <div>
      <CodeEditor value={content} readonly language="python" />
    </div>
  </Dialog.Content>
</Dialog.Root>
