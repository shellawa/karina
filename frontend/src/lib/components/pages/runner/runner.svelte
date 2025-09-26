<script module lang="ts">
  import { Badge } from "$lib/components/ui/badge"
  import { Button } from "$lib/components/ui/button"
  import * as Card from "$lib/components/ui/card"
  import { Label } from "$lib/components/ui/label"
  import * as Select from "$lib/components/ui/select"
  import {
    GetProblems,
    GetParticipants,
    GetTestCases,
    GetSolveResults
  } from "$lib/wailsjs/go/models/Service"
  import { RunAllParticipants } from "$lib/wailsjs/go/languages/Service"
  import { EventsOn } from "$lib/wailsjs/runtime/runtime"
  import { ChartBar, Code, Download, Play, User } from "@lucide/svelte"
  import AddAndEditDialog from "./addAndEditDialog.svelte"
  import AddParticipantDialog from "./addParticipantDialog.svelte"
  import Progress from "$lib/components/ui/progress/progress.svelte"
  import { flip } from "svelte/animate"
  import type { models } from "$lib/wailsjs/go/models"

  let problems = $state(await GetProblems())

  // svelte-ignore state_referenced_locally
  let selectedProblemId = $state(problems[0]?.id ?? "")
  let selectedProblem = $derived(problems.find((x) => x.id == selectedProblemId))

  let participantsPromise = $derived(GetParticipants(selectedProblemId))

  EventsOn("problem:change", async (id) => {
    problems = await GetProblems()
    selectedProblemId = id
    solveStatus = "idle"
  })

  EventsOn("participant:change", async () => {
    participantsPromise = GetParticipants(selectedProblemId)
  })

  let solveProgress = $state({
    participantId: "",
    language: "",
    testNo: 0
  })

  // idle || running || finished
  let solveStatus = $state("idle")
  EventsOn("solve:run_status", async (status) => {
    solveStatus = status
  })

  // svelte-ignore state_referenced_locally
  let solveResults = $state([]) as models.ParticipantSolveResult[]
  EventsOn("solve:test_run_start", async (solveStatusUpdate) => {
    solveProgress = solveStatusUpdate
  })

  EventsOn("solve:test_run_finish", async () => {
    solveResults = await GetSolveResults(selectedProblemId)
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
          <Select.Root
            type="single"
            bind:value={selectedProblemId}
            onValueChange={() => {
              solveStatus = "idle"
            }}
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
    {#await Promise.all( [participantsPromise, GetTestCases(selectedProblemId)] ) then [participants, testCases]}
      <Card.Root>
        <Card.Header>
          <Card.Title class="flex items-center gap-x-2">
            <ChartBar class="h-5 w-5" />
            Execution Status
          </Card.Title>
        </Card.Header>
        <Card.Content class="space-y-3">
          <div class="rounded-lg border-l-4 border-blue-500 bg-blue-50 p-3">
            {#if solveStatus == "idle"}
              <div class="text-sm font-semibold text-blue-900">Idle</div>
            {:else if solveStatus == "running"}
              <div class="text-sm font-semibold text-blue-900">Currently Running</div>
            {:else if solveStatus == "finished"}
              <div class="text-sm font-semibold text-blue-900">Finished</div>
            {/if}

            {#if solveStatus != "idle"}
              <div class="mt-1 text-sm text-blue-700">
                {participants.find((x) => x.id == solveProgress.participantId)?.name} ({solveProgress.language})
              </div>
              <div class="mt-1 text-xs font-medium text-blue-600">
                Test {solveProgress.testNo}/{testCases.length}
              </div>
            {/if}
          </div>

          <div class="space-y-1.5">
            <div class="text-sm font-semibold">Queue Status</div>
            <div class="space-y-1 text-sm">
              <div class="flex justify-between">
                <span class="font-medium text-green-600">Completed</span>
                <span class="font-mono font-semibold">3/6</span>
              </div>
              <div class="flex justify-between">
                <span class="font-medium text-blue-600">Running</span>
                <span class="font-mono font-semibold">1/6</span>
              </div>
              <div class="flex justify-between">
                <span class="font-medium text-gray-600">Pending</span>
                <span class="font-mono font-semibold">2/6</span>
              </div>
            </div>
          </div>
          <Progress
            value={participants.findIndex((x) => x.id == solveProgress.participantId) *
              testCases.length +
              solveProgress.testNo}
            max={participants.length * testCases.length}
          />
        </Card.Content>
      </Card.Root>
    {/await}
  </div>

  {#await Promise.all( [participantsPromise, GetTestCases(selectedProblemId)] ) then [participants, testCases]}
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
          {#each participants
            .map((p) => {
              const res = solveResults.find((x) => x.id === p.id)?.results ?? []
              const ac = res.filter((z) => z.verdict === "AC")

              const acCount = ac.length
              const totalTime = ac.reduce((s, z) => s + z.time, 0)
              const totalMemory = ac.reduce((s, z) => s + z.memory, 0)

              const minTime = ac.length ? Math.min(...ac.map((z) => z.time)) : 0
              const maxTime = ac.length ? Math.max(...ac.map((z) => z.time)) : 0
              const minMemory = ac.length ? Math.min(...ac.map((z) => z.memory)) : 0
              const maxMemory = ac.length ? Math.max(...ac.map((z) => z.memory)) : 0

              return { ...p, acCount, totalTime, totalMemory, minTime, maxTime, minMemory, maxMemory }
            })
            .sort((a, b) => {
              if (a.acCount !== b.acCount) return b.acCount - a.acCount
              if (a.totalTime !== b.totalTime) return a.totalTime - b.totalTime
              return a.totalMemory - b.totalMemory
            }) as ptcp, index (ptcp.id)}
            <div
              animate:flip
              class="grid-cols-22 grid gap-4 rounded-lg border px-4 py-3 transition-colors hover:bg-gray-50"
            >
              <!-- rank -->
              <div class="col-span-2 flex items-center font-medium text-gray-600">#{index + 1}</div>

              <!-- name and organization -->
              <div class="col-span-6">
                <div class="font-medium">{ptcp.name}</div>
                <div class="text-xs text-gray-500">
                  {ptcp.id} &bull; {ptcp.organization}
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
                  <div class="text-xs font-medium text-gray-700">
                    {ptcp.acCount}/{testCases.length}
                  </div>
                </div>
              </div>

              <!-- avg time and time range -->
              <div class="col-span-3 flex items-center font-mono">
                <div class="flex flex-col items-center">
                  <div class="">{ptcp.totalTime}ms</div>
                  <div class="text-xs text-gray-500">{ptcp.minTime}-{ptcp.maxTime}</div>
                </div>
              </div>

              <!-- avg memory and memory range -->
              <div class="col-span-3 flex items-center font-mono">
                <div class="flex flex-col items-center">
                  <div class="">{ptcp.totalMemory}MB</div>
                  <div class="text-xs text-gray-500">{ptcp.minMemory}-{ptcp.maxMemory}</div>
                </div>
              </div>
            </div>
          {/each}
        </div>
      </Card.Content>
    </Card.Root>
  {/await}
</div>
