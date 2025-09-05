<script>
  import { Button } from "$lib/components/ui/button/index.js"
  import { Moon, Timer, Settings } from "@lucide/svelte"
</script>

<header class="border-b bg-white shadow-sm">
  <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
    <div class="flex h-16 items-center justify-between">
      <div class="flex items-center space-x-3">
        <Timer class="h-8 w-8 text-blue-600" />
        <h1 class="text-xl font-bold">karina</h1>
      </div>

      <div class="flex items-center space-x-3">
        <Button variant="outline" size="sm">
          <Settings />
          Settings
        </Button>

        <Button variant="outline" size="icon" class="size-8">
          <Moon />
        </Button>
      </div>
    </div>
  </div>
</header>
