<script lang="ts">
  import { Button, buttonVariants } from "$lib/components/ui/button"
  import * as Dialog from "$lib/components/ui/dialog"
  import Input from "$lib/components/ui/input/input.svelte"
  import { Label } from "$lib/components/ui/label"
  import * as Select from "$lib/components/ui/select"
  import { Textarea } from "$lib/components/ui/textarea"
  import { type models } from "$lib/wailsjs/go/models"
  import { WriteProblem } from "$lib/wailsjs/go/models/Service"
  import { Plus, Settings, Trash } from "@lucide/svelte"

  let { dialogType, problem }: { dialogType: string; problem: models.Problem } = $props()
  let fieldProblem = $state(JSON.parse(JSON.stringify(problem ?? {}))) as models.Problem

  $inspect(fieldProblem)
</script>

<Dialog.Root>
  <Dialog.Trigger
    class={buttonVariants({ variant: "outline", class: "h-10" })}
    disabled={dialogType == "edit" && !problem}
  >
    {#if dialogType == "add"}
      <Plus /> Add
    {:else if dialogType == "edit"}
      <Settings /> Edit
    {/if}
  </Dialog.Trigger>

  <Dialog.Content class="lg:max-w-4xl">
    <Dialog.Header>
      <Dialog.Title>
        {#if dialogType == "add"}
          Add a problem
        {:else if dialogType == "edit"}
          Edit current problem
        {/if}
      </Dialog.Title>
    </Dialog.Header>

    <div class="grid grid-cols-1 gap-4 lg:grid-cols-2">
      <div class="space-y-3">
        <Label class="mb-2">ID</Label>
        <Input
          type="text"
          placeholder="two_sum"
          bind:value={fieldProblem.id}
          disabled={dialogType == "edit"}
        />

        <Label class="mb-2">Label</Label>
        <Input type="text" placeholder="Two sum" bind:value={fieldProblem.label} />

        <Label class="mb-2">Description</Label>
        <Textarea
          placeholder="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
          bind:value={fieldProblem.description}
        />

        <Label class="mb-2">IO Mode</Label>
        <Select.Root
          type="single"
          bind:value={
            () => fieldProblem.io_mode.toString(), (v: string) => (fieldProblem.io_mode = Number(v))
          }
        >
          <Select.Trigger class="w-[180px]">
            {fieldProblem.io_mode ? "File IO" : "Standard IO"}
          </Select.Trigger>
          <Select.Content>
            <Select.Item value="0">Standard IO</Select.Item>
            <Select.Item value="1">File IO</Select.Item>
          </Select.Content>
        </Select.Root>
      </div>

      <div class="space-y-3">
        <Label class="">Default Limits</Label>
        <div class="grid grid-cols-2 gap-3">
          <div>
            <Label class="mb-2">Time (ms)</Label>
            <Input type="number" bind:value={fieldProblem.time} />
          </div>
          <div>
            <Label class="mb-2">Memory (MB)</Label>
            <Input type="number" bind:value={fieldProblem.memory} />
          </div>
        </div>
      </div>
    </div>

    <Dialog.Footer class="border-t pt-5">
      {#if dialogType == "add"}
        <Dialog.Close class="float-right">
          <Button
            variant="default"
            onclick={() => {
              WriteProblem(fieldProblem)
            }}>Add problem</Button
          >
        </Dialog.Close>
      {:else if dialogType == "edit"}
        <div class="flex w-full items-stretch justify-between">
          <Dialog.Close>
            <Button variant="destructive"><Trash /> Delete</Button>
          </Dialog.Close>
          <Dialog.Close>
            <Button
              variant="default"
              onclick={() => {
                WriteProblem(fieldProblem)
              }}>Save changes</Button
            >
          </Dialog.Close>
        </div>
      {/if}
    </Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>
