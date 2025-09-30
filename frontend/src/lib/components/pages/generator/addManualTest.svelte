<script lang="ts">
  import * as Dialog from "$lib/components/ui/dialog"
  import { Plus } from "@lucide/svelte"
  import { Button, buttonVariants } from "$lib/components/ui/button"
</script>

<Dialog.Root>
  <Dialog.Trigger disabled class={buttonVariants({ variant: "outline", size: "sm" })}>
    <Plus /> Add manually
  </Dialog.Trigger>
</Dialog.Root>
