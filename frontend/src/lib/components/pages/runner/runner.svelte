<script lang="ts">
  import * as Card from "$lib/components/ui/card/index.js"
  import * as Select from "$lib/components/ui/select/index.js"
  import { Label } from "$lib/components/ui/label/index.js"
  import { Code, Download, Play, Plus, Settings } from "@lucide/svelte"
  import { Button } from "$lib/components/ui/button/index.js"

  const problems = [
    {
      slug: "two_sum",
      label: "Two Sum",
      description:
        "lorem ipsumLorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam."
    },
    {
      slug: "three_sum",
      label: "Three Sum",
      description:
        "lorem ipsumLorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam."
    }
  ]
  let selectedProblemSlug = $state(problems[0].slug)
  let selectedProblem = $derived(problems.find((x) => x.slug == selectedProblemSlug))
</script>

<div class="grid grid-cols-1 gap-4 lg:grid-cols-4">
  <Card.Root class="lg:col-span-1">
    <Card.Header>
      <Card.Title class="flex items-center gap-x-2">
        <Code class="h-5 w-5" />
        <span>Problem Configuration </span>
      </Card.Title>
    </Card.Header>

    <Card.Content class="space-y-3">
      <div>
        <Label>Select problem</Label>
        <Select.Root type="single" bind:value={selectedProblemSlug}>
          <Select.Trigger>
            {selectedProblem?.label}
          </Select.Trigger>

          <Select.Content>
            {#each problems as problem}
              <Select.Item value={problem.slug}>{problem.label}</Select.Item>
            {/each}
          </Select.Content>
        </Select.Root>
      </div>

      <div class="rounded-lg bg-gray-50 p-3">
        <div class="line-clamp-4 text-justify text-xs text-gray-600">
          {selectedProblem?.description}
        </div>
        <div class="line-clamp-4 text-xs text-gray-500">
          {Math.floor(Math.random() * 10)} test cases
        </div>
      </div>

      <Button class="h-12 w-full">
        <Play /> Run all participants
      </Button>
      <div class="grid grid-cols-2 gap-2">
        <Button variant="outline" class="h-10">
          <Plus /> Add
        </Button>

        <Button variant="outline" class="h-10">
          <Settings /> Edit
        </Button>
      </div>

      <Button variant="outline" class="h-10 w-full">
        <Download /> Export Results
      </Button>
    </Card.Content>
  </Card.Root>
</div>
