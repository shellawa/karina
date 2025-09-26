<script module lang="ts">
  import { Badge } from "$lib/components/ui/badge"
  import { Button } from "$lib/components/ui/button"
  import * as Card from "$lib/components/ui/card"
  import { Label } from "$lib/components/ui/label"
  import * as Select from "$lib/components/ui/select"
  import { GetProblems, GetParticipants, GetTestCases } from "$lib/wailsjs/go/models/Service"
  import { RunAllParticipants } from "$lib/wailsjs/go/languages/Service"
  import { EventsOn } from "$lib/wailsjs/runtime/runtime"
  import { ChartBar, Code, Download, Play, User } from "@lucide/svelte"
  import AddAndEditDialog from "./addAndEditDialog.svelte"
  import AddParticipantDialog from "./addParticipantDialog.svelte"

  let problems = $state(await GetProblems())

  // svelte-ignore state_referenced_locally
  let selectedProblemId = $state(problems[0]?.id ?? "")
  let selectedProblem = $derived(problems.find((x) => x.id == selectedProblemId))

  let participantsPromise = $derived(GetParticipants(selectedProblemId))

  EventsOn("problem:change", async (id) => {
    problems = await GetProblems()
    selectedProblemId = id
  })

  EventsOn("participant:change", async () => {
    participantsPromise = GetParticipants(selectedProblemId)
  })

  EventsOn("run:python:sample", async (data) => {
    console.log(data)
  })
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
          <Select.Root type="single" bind:value={selectedProblemId}>
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

        <div class="rounded-lg bg-gray-50 p-3">
          <div class="line-clamp-4 text-justify text-xs text-gray-600">
            {selectedProblem?.description}
          </div>
          <div class="line-clamp-4 text-xs text-gray-500">
            {#await GetTestCases(selectedProblemId) then testCases}
              {testCases.length}
            {/await} test case(s)
          </div>
        </div>

        <Button
          class="h-12 w-full"
          onclick={() => RunAllParticipants(selectedProblemId, 1500, 1024)}
        >
          <Play /> Run all participants
        </Button>
        <div class="grid grid-cols-2 gap-2">
          {#key selectedProblemId}
            <AddAndEditDialog
              dialogType="add"
              problem={{ id: "", description: "", label: "", time: 1500, memory: 512 }}
            />

            <AddAndEditDialog dialogType="edit" problem={selectedProblem!} />
          {/key}
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
      <Card.Title class="flex items-center justify-between">
        <div class="flex items-center gap-x-2">
          <User class="h-5 w-5" />
          Leaderboard
        </div>
        {#key participantsPromise}
          <AddParticipantDialog
            participant={{ id: "", name: "", organization: "" }}
            problemId={selectedProblemId}
          />
        {/key}
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
        {#await participantsPromise then participants}
          {#each participants as participant, index}
            <div
              class="grid-cols-22 grid gap-4 rounded-lg border px-4 py-3 transition-colors hover:bg-gray-50"
            >
              <!-- rank -->
              <div class="col-span-2 flex items-center font-medium text-gray-600">#{index + 1}</div>

              <!-- name and organization -->
              <div class="col-span-6">
                <div class="font-medium">{participant.name}</div>
                <div class="text-xs text-gray-500">
                  {participant.id} &bull; {participant.organization}
                </div>
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
          {/each}
        {/await}
      </div>
    </Card.Content>
  </Card.Root>
</div>
