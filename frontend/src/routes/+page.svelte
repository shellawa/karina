<script>
  import Header from "$lib/components/layout/header.svelte"
  import Runner from "$lib/components/pages/runner/runner.svelte"
  import Generator from "$lib/components/pages/generator/generator.svelte"

  import * as Tabs from "$lib/components/ui/tabs/index.js"

  import { Play, FlaskConical } from "@lucide/svelte"

  let tab = $state("runner")
</script>

<Header />

<main class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
  <Tabs.Root value="runner" onValueChange={(v) => (tab = v)}>
    <Tabs.List class="w-full">
      <Tabs.Trigger value="runner"><Play /> Runner</Tabs.Trigger>
      <Tabs.Trigger value="generator"><FlaskConical /> Generator</Tabs.Trigger>
    </Tabs.List>
    <Tabs.Content value="runner">
      {#key tab}<Runner />{/key}
    </Tabs.Content>
    <Tabs.Content value="generator">
      {#key tab}<Generator />{/key}
    </Tabs.Content>
  </Tabs.Root>
</main>
