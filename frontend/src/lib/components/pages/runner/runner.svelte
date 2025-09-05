<script lang="ts">
  import * as Card from "$lib/components/ui/card/index.js"
  import * as Select from "$lib/components/ui/select/index.js"
  import { Button } from "$lib/components/ui/button/index.js"
  import { Label } from "$lib/components/ui/label/index.js"
  import { Badge } from "$lib/components/ui/badge/index.js"
  import { ChartBar, Code, Download, Play, Plus, Settings, User } from "@lucide/svelte"

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
  <div class="space-y-4 lg:col-span-1">
    <Card.Root>
      <Card.Header>
        <Card.Title class="flex items-center gap-x-2">
          <Code class="h-5 w-5" />
          <span>Problem Configuration </span>
        </Card.Title>
      </Card.Header>

      <Card.Content class="space-y-3">
        <div>
          <Label class="mb-2">Select problem</Label>
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

    <Card.Root>
      <Card.Header>
        <Card.Title class="flex items-center gap-x-2">
          <ChartBar class="h-5 w-5" />
          Execution Status
        </Card.Title>
      </Card.Header>
    </Card.Root>
  </div>

  <Card.Root class="lg:col-span-3">
    <Card.Header>
      <Card.Title class="flex items-center gap-x-2">
        <User class="h-5 w-5" />
        Leaderboard
      </Card.Title>
    </Card.Header>
    <Card.Content>
      <!-- table header -->
      <div
        class="grid-cols-22 mb-4 grid gap-4 rounded-lg bg-gray-50 px-4 py-3 text-sm font-medium text-gray-600"
      >
        <div class="col-span-2">#</div>
        <div class="col-span-6">Name</div>
        <div class="col-span-4">Language</div>
        <div class="col-span-4">Status</div>
        <div class="col-span-3">Time</div>
        <div class="col-span-3">Memory</div>
      </div>

      <!-- table content -->
      <div class="space-y-2">
        <div
          class="grid-cols-22 grid gap-4 rounded-lg border px-4 py-3 transition-colors hover:bg-gray-50"
        >
          <!-- rank -->
          <div class="col-span-2 flex items-center font-medium text-gray-600">#1</div>

          <!-- name and organization -->
          <div class="col-span-6">
            <div class="font-medium">Alice Johnson</div>
            <div class="text-xs text-gray-500">United Nations</div>
          </div>

          <!-- language badge -->
          <div class="col-span-4 flex items-center">
            <Badge variant="outline">Python</Badge>
          </div>

          <!-- status badge and passed/total -->
          <div class="col-span-4 flex items-center">
            <div class="flex flex-col items-center">
              <Badge class="bg-green-400">Accepted</Badge>
              <div class="text-xs text-gray-500">10/10</div>
            </div>
          </div>

          <!-- avg time and time range -->
          <div class="col-span-3 flex items-center font-mono">
            <div class="flex flex-col items-center">
              <div class="">123ms</div>
              <div class="text-xs text-gray-500">54-612</div>
            </div>
          </div>

          <!-- avg memory and memory range -->
          <div class="col-span-3 flex items-center font-mono">
            <div class="flex flex-col items-center">
              <div class="">64MB</div>
              <div class="text-xs text-gray-500">21-727</div>
            </div>
          </div>
        </div>
      </div>
    </Card.Content>
  </Card.Root>
</div>
