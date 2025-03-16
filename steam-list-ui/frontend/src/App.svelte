<script lang="ts">
    import { onMount } from "svelte";
    import type { Game } from "./model/game";
    import { getPlayerGames } from "./service/games";

    let steamId = "";
    let steamIdInput = "";
    let games: Game[] = [];
    let page = 1;
    let hasGames = true;
    let fetchinGames = false;

    function handleScroll() {
        const bottom =
            window.innerHeight + window.scrollY >=
            document.body.offsetHeight + 400;
        if (bottom && !fetchinGames && hasGames) {
            getGames();
        }
    }

    onMount(() => {
        window.addEventListener("scroll", handleScroll);
        return () => {
            window.removeEventListener("scroll", handleScroll);
        };
    });

    function startSync() {
        console.log(steamId);
        games = [];
        steamId = steamIdInput;
        getGames();
    }

    async function getGames() {
        if (steamId == "") return;
        fetchinGames = true;
        let gamesFetch: Game[] = await getPlayerGames(steamId.trim(), page);
        if (gamesFetch === null) {
            hasGames = false;
            fetchinGames = false;
            return;
        }
        page++;
        games =
            games.concat(gamesFetch).filter((item, index, array) => {
                return (
                    array.findIndex((i) => i.igdbid === item.igdbid) === index
                );
            }) ?? [];
        fetchinGames = false;
    }
</script>

<main class="bg-black">
    <div>
        <div class="container p-4">
            <div class="row flex-row justify-content-center align-items-center">
                <div class="col-5 flex-row d-flex gap-2">
                    <div class="form-group flex-grow-1">
                        <input
                            bind:value={steamIdInput}
                            class="form-control"
                            id="steamid"
                            placeholder="Steam ID"
                        />
                    </div>

                    <button
                        class="btn btn-outline-primary"
                        on:click={startSync}
                    >
                        Sync</button
                    >
                </div>
            </div>
        </div>
        {#if games.length == 0}
            <div></div>
        {:else}
            <div class="container">
                <div class="row">
                    {#each games as game}
                        <div class="col-3 pt-3 pe-2">
                            <img
                                src={game.cover_hd_url}
                                alt=""
                                class="rounded-bottom-5"
                            />
                            <div class="text-white pt-2">{game.name}</div>
                        </div>
                    {/each}
                </div>
            </div>
        {/if}
        {#if fetchinGames}
            <div class="h1 text-white">LOADING...</div>
        {:else}
            <div></div>
        {/if}
    </div>
</main>
