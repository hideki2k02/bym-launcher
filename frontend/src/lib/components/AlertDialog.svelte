<script>
    export let open = false
    export let error = ""

    import { Button, buttonVariants } from "$lib/components/ui/button";
    import {Quit} from "$wails/runtime/runtime.js";
    import {
        Dialog,
        DialogContent, DialogDescription,
        DialogFooter,
        DialogHeader,
        DialogTitle,
        DialogOverlay,
        DialogTrigger
    } from "$lib/components/ui/dialog";
    import {XIcon} from "lucide-svelte";
</script>
<Dialog bind:open={open}>
<!--    <DialogTrigger class={buttonVariants({ variant: "ghost" })}>-->
<!--        <XIcon class="h-4 w-4"></XIcon>-->
<!--    </DialogTrigger>-->
<!--    <DialogOverlay class="bg-red-500"></DialogOverlay>-->
    <DialogContent class="text-left bg-background text-foreground">
        <DialogHeader class="text-left">
            <DialogTitle class="select-none">Oopsie</DialogTitle>
            <DialogDescription>
                <p class="text-secondary-foreground mt-4 mb-4">
                    Launcher caught an error: <b class="font-bold capitalize">{error}</b>
                </p>
            </DialogDescription>
        </DialogHeader>
        <DialogFooter>
            <div class="flex justify-end gap-2">
                <Button variant="secondary" type="button" on:click={() => open = false}>Continue</Button>
                <Button variant="secondary" type="button" on:click={Quit}>Quit</Button>
            </div>
        </DialogFooter>
    </DialogContent>
</Dialog>