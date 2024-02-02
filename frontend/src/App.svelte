<script>
  import { LaunchGame, InitializeApp } from '$wails/go/main/App.js';
  import { Button } from '$lib/components/ui/button';
  import * as Select from '$lib/components/ui/select';
  import { Loader2 } from 'lucide-svelte';
  import { setMode } from 'mode-watcher';
  import LightDarkButton from '$lib/components/LightDarkButton.svelte';
  import AlertDialog from '$lib/components/AlertDialog.svelte';
  import { Quit, EventsOn, LogPrint } from '$wails/runtime/runtime.js';

  // import { EventsOn } from '../../wailsjs/runtime'

  setMode('dark');

  const builds = [
    { value: 'stable', label: 'Stable' },
    { value: 'http', label: 'Http' },
    { value: 'local', label: 'Local' },
  ];
  let build = builds[0];

  const runtimes = [
    { value: 'flashplayer', label: 'Flash Player' },
    { value: 'ruffle', label: 'Ruffle Player (Experimental)' },
  ];

  let runtime = runtimes[0];

  let disabled = false;
  let showError = false;
  let errorCode = '';

  let debugLogs = [];

  EventsOn('infoLog', (event) => {
    debugLogs = [...debugLogs, event];
  });

  InitializeApp().then(() => {
    debugLogs = [...debugLogs, 'Launcher initialized'];
  });

  function launch() {
    disabled = true;
    LaunchGame(build.value, runtime.value)
      .then(() => {
        showError = false;
        Quit();
      })
      .catch((err) => {
        errorCode = err;
        showError = true;
      });
    disabled = false;
  }
</script>

<main
  class="h-screen bg-background text-foreground flex flex-col gap-4 p-4 antialiased select-none font-sans"
>
  <div class="flex justify-between">
    <h1 class="text-foreground font-bold text-2xl">
      Backyard Monsters <br />
      Refitted
    </h1>
    <LightDarkButton />
  </div>
  <div class="grow bg-secondary font-mono p-3">
    <h5>Welcome to Bym Refitted!</h5>
    {#each debugLogs as log}
      <p><small>{log}</small></p>
    {/each}

    <!-- <p><small>Server status: Online - v0.2.3-alpha</small></p>
    <p class="text-red"><small>Current version: v0.2.1-alpha</small></p>
    <p class=""><small>Installing v0.2.3-alpha....</small></p>
    <p><small class="text-green">Complete</small></p>
    <p class="text-red"><small>Flash player detected</small></p> -->
  </div>
  <div class="mt-auto w-full flex justify-between">
    <label>SWF Build</label>
    <Select.Root bind:selected={build} portal={null}>
      <Select.Trigger class="w-[180px] rounded">
        <Select.Value placeholder="Select Build" />
      </Select.Trigger>
      <Select.Content>
        <Select.Group>
          {#each builds as build}
            <Select.Item value={build.value} label={build.label}>{build.label}</Select.Item>
          {/each}
        </Select.Group>
      </Select.Content>
      <Select.Input name="build" />
    </Select.Root>
  </div>
  <div class="mt-auto w-full flex justify-between">
    <label>Flash Runtime</label>
    <Select.Root bind:selected={runtime} portal={null}>
      <Select.Trigger class="w-[180px] rounded">
        <Select.Value placeholder="Select Build" />
      </Select.Trigger>
      <Select.Content>
        <Select.Group>
          {#each runtimes as runtime}
            <Select.Item value={runtime.value} label={runtime.label}>{runtime.label}</Select.Item>
          {/each}
        </Select.Group>
      </Select.Content>
      <Select.Input name="runtime" />
    </Select.Root>
  </div>

  <div class="mt-auto w-full flex justify-between">
    <Button variant="outline" class="p-4 rounded w-32" on:click={launch} {disabled}>
      {#if disabled}
        <Loader2 />
      {:else}
        Launch Game
      {/if}
    </Button>
  </div>

  <AlertDialog bind:open={showError} error={errorCode}></AlertDialog>
</main>
