<script lang="ts">
  import { Badge } from "$lib/components/ui/badge"
  import { buttonVariants } from "$lib/components/ui/button"
  import * as Dialog from "$lib/components/ui/dialog"
  import type { ParticipantSolveDetails } from "$lib/types"
  import { List } from "@lucide/svelte"
  import CodeViewDialog from "./codeViewDialog.svelte"
  import TestDetailsDialog from "./testDetailsDialog.svelte"

  let { details }: { details: ParticipantSolveDetails } = $props()
</script>

<Dialog.Root>
  <Dialog.Trigger
    class={buttonVariants({ variant: "outline", size: "sm" })}
    disabled={details.status == "Running" || details.status == "Pending"}
  >
    <List />
  </Dialog.Trigger>
  <Dialog.Content class="lg:max-w-5xl">
    <Dialog.Header>
      <Dialog.Title>{details.name}</Dialog.Title>
    </Dialog.Header>

    <div class="bg-muted/30 grid grid-cols-6 gap-4 rounded-lg p-4 text-center">
      <div class="space-y-1">
        <div class="text-2xl font-bold text-green-600">
          {details.acCount}
        </div>
        <div class="text-muted-foreground text-sm font-semibold">Passed</div>
      </div>

      <div class="space-y-1">
        <div class="text-2xl font-bold text-red-600">
          {details.testCases.length - details.acCount}
        </div>
        <div class="text-muted-foreground text-sm font-semibold">Failed</div>
      </div>

      <div class="space-y-1">
        <div class="text-2xl font-bold text-gray-900">
          {(details.totalTime / details.acCount).toFixed(0)}ms
        </div>
        <div class="text-muted-foreground text-sm font-semibold">Avg Time</div>
      </div>

      <div class="space-y-1">
        <div class="text-2xl font-bold text-gray-900">
          {details.minTime}-{details.maxTime}ms
        </div>
        <div class="text-muted-foreground text-sm font-semibold">Time Range</div>
      </div>

      <div class="space-y-1">
        <div class="text-2xl font-bold text-gray-900">
          {(details.totalMemory / details.acCount).toFixed(0)}MB
        </div>
        <div class="text-muted-foreground text-sm font-semibold">Avg Memory</div>
      </div>

      <div class="space-y-1">
        <div class="text-2xl font-bold text-gray-900">
          {details.minMemory}-{details.maxMemory}MB
        </div>
        <div class="text-muted-foreground text-sm font-semibold">Memory Range</div>
      </div>
    </div>

    <div class="grid max-h-[400px] grid-cols-1 gap-3 overflow-y-auto md:grid-cols-2">
      {#each details.res as solve, index}
        <Dialog.Root>
          <Dialog.Trigger>
            <div class="hover:bg-muted/20 cursor-pointer rounded-lg border p-3 transition-colors">
              <div class="mb-2 flex items-center justify-between">
                <span class="font-semibold text-gray-900">{details.testCases[index].Id}</span>
                <Badge
                  class={`${
                    solve.verdict === "AC"
                      ? "bg-green-500"
                      : solve.verdict === "WA"
                        ? "bg-red-500"
                        : solve.verdict === "TLE"
                          ? "bg-yellow-500"
                          : "bg-purple-500"
                  } px-2 py-1 text-xs font-semibold text-white`}
                >
                  {solve.verdict}
                </Badge>
              </div>
              <div class="grid grid-cols-2 gap-3 text-sm">
                <div>
                  <div class="text-muted-foreground font-medium">Time</div>
                  <div class="font-mono font-semibold text-gray-900">{solve.time}ms</div>
                </div>
                <div>
                  <div class="text-muted-foreground font-medium">Memory</div>
                  <div class="font-mono font-semibold text-gray-900">{solve.memory}MB</div>
                </div>
              </div>
            </div>
          </Dialog.Trigger>
          <TestDetailsDialog testCase={details.testCases[index]} {solve} />
        </Dialog.Root>
      {/each}
    </div>

    <div class="max-w-fit">
      <CodeViewDialog content={details.code ?? ""} />
    </div>
  </Dialog.Content>
</Dialog.Root>
