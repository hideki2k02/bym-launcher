<script>
  import {Greet, LaunchGame} from '$wails/go/main/App.js'
  import {Button} from "$lib/components/ui/button";
  import * as Select from "$lib/components/ui/select";
  import {Loader2} from "lucide-svelte";
  import { ModeWatcher, setMode } from "mode-watcher";
  import LightDarkButton from "$lib/components/LightDarkButton.svelte";
  import AlertDialog from "$lib/components/AlertDialog.svelte";
  import {Quit} from "$wails/runtime/runtime.js";

  setMode('dark')

  let resultText = "Please enter your name below 👇"
  let build = { value: "stable", label: "Stable" }

  const builds = [
    { value: "stable", label: "Stable" },
    { value: "http", label: "Http" },
    { value: "local", label: "Local" },
  ];

  let disabled = false
  let showError = false
  let errorCode = ""

  function launch() {
    disabled = true
    LaunchGame(build.value)
            .then(() => {
              showError = false
              Quit();
            })
            .catch(err => {
              errorCode = err;
              // if (err.includes("file does not exist")) { errorCode = "0x0419" } else {  errorCode = "0x556e6b6e6f776e" }
              showError = true
            })
    disabled = false
  }
</script>

<main class="h-screen bg-background text-foreground flex flex-col gap-4 p-4 antialiased select-none font-sans">
  <div class="flex justify-between">
    <h1 class="text-foreground font-bold text-2xl">Backyard Monsters <br> Refitted</h1>
    <LightDarkButton/>
  </div>
  <div class="grow bg-secondary"></div>
  <div class="mt-auto w-full flex justify-between">
    <Select.Root bind:selected="{build}" portal={null}>
      <Select.Trigger  class="w-[180px] rounded">
        <Select.Value placeholder="Select Build" />
      </Select.Trigger>
      <Select.Content>
        <Select.Group>
          {#each builds as build}
            <Select.Item value={build.value} label={build.label}
            >{build.label}</Select.Item
            >
          {/each}
        </Select.Group>
      </Select.Content>
      <Select.Input name="favoriteFruit" />
    </Select.Root>
    <Button variant="outline" class="p-4 rounded w-32" on:click={launch} {disabled}>
        {#if (disabled)}
            <Loader2/>
            {:else }
            Launch Game
        {/if}
    </Button>
  </div>

  <AlertDialog bind:open={showError} error="{errorCode}"></AlertDialog>
</main>
