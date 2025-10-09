<script lang="ts">
  import * as Dialog from "$lib/components/ui/dialog"
  import type { common, models } from "$lib/wailsjs/go/models"
  let { testCase, solve }: { testCase: models.TestCase; solve: common.TestSolveResult } = $props()
</script>

<Dialog.Content class="lg:max-w-4xl">
  <div class="mb-6 flex items-center justify-between">
    <div>
      <h2 class="text-xl font-bold text-gray-900">
        {testCase.Id}
      </h2>
      <p class="text-muted-foreground mt-1 text-sm font-medium">
        Verdict: {solve.verdict} • {solve.time}ms • {solve.memory}MB
      </p>
    </div>
  </div>

  <div class="space-y-6 overflow-auto">
    <div class="max-h-[200px] max-w-full">
      <h3 class="mb-3 text-lg font-semibold text-gray-900">Input</h3>
      <div class="overflow-auto rounded-lg border bg-gray-50 p-4">
        <pre class="whitespace-pre font-mono">{testCase.Input.trimEnd() || "Not available"}</pre>
      </div>
    </div>

    <div class="max-h-[200px] max-w-full">
      <h3 class="mb-3 text-lg font-semibold text-gray-900">Expected Output</h3>
      <div class="overflow-auto rounded-lg border border-green-200 bg-green-50 p-4">
        <pre class="whitespace-pre font-mono">{testCase.Output.trimEnd() || "Not available"}</pre>
      </div>
    </div>

    <div class="max-h-[200px] max-w-full">
      <h3 class="mb-3 text-lg font-semibold text-gray-900">Actual Output</h3>
      <div
        class={"overflow-auto rounded-lg border p-4 " +
          (solve.verdict === "AC" ? "border-green-200 bg-green-50" : "border-red-200 bg-red-50")}
      >
        <pre class="whitespace-pre font-mono">{solve.actual_output.trimEnd() ||
            "Not available"}</pre>
      </div>
    </div>
  </div>
</Dialog.Content>
