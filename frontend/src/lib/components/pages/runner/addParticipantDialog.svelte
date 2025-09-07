<script lang="ts">
  import * as Dialog from "$lib/components/ui/dialog"
  import Input from "$lib/components/ui/input/input.svelte"
  import { Label } from "$lib/components/ui/label"
  import { Button, buttonVariants } from "$lib/components/ui/button"
  import { Plus } from "@lucide/svelte"
  import { AddOneParticipant } from "$lib/wailsjs/go/models/ParticipantService"
  import { type models } from "$lib/wailsjs/go/models"

  let { participant, problemId }: { participant: models.Participant; problemId: string } = $props()
  let fieldParticipant = $state(JSON.parse(JSON.stringify(participant ?? {}))) as models.Participant
</script>

<Dialog.Root>
  <Dialog.Trigger class={buttonVariants({ variant: "outline", size: "sm" })}>
    <Plus /> Add participants
  </Dialog.Trigger>
  <Dialog.Content>
    <Dialog.Header>
      <Dialog.Title>Add a participant</Dialog.Title>
    </Dialog.Header>

    <form>
      <div class="space-y-3">
        <Label class="mb-2">ID</Label>
        <Input type="text" placeholder="john_smith" bind:value={fieldParticipant.id} />
        <Label class="mb-2">Name</Label>
        <Input type="text" placeholder="John Smith" bind:value={fieldParticipant.name} />
        <Label class="mb-2">Organization</Label>
        <Input
          type="text"
          placeholder="United Nations"
          bind:value={fieldParticipant.organization}
        />
      </div>
    </form>
    <Dialog.Footer class="border-t pt-5">
      <Dialog.Close class="float-right">
        <Button
          type="submit"
          onclick={async () => await AddOneParticipant(fieldParticipant, problemId)}
        >
          Add
        </Button>
      </Dialog.Close>
    </Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>
