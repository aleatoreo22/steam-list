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
        <div class="container p-4 mx-auto">
            <div class="flex flex-row justify-center items-center">
                <div class="flex flex-row gap-2 w-full max-w-md">
                    <div class="flex-grow">
                        <input
                            bind:value={steamIdInput}
                            class="text-white w-full px-3 py-2 border border-gray-300 rounded-md transition focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                            id="steamid"
                            placeholder="Steam ID"
                        />
                    </div>

                    <button
                        class="px-4 py-2 border border-blue-500 text-blue-500 rounded-md hover:bg-blue-500 hover:text-white transition"
                        on:click={startSync}
                    >
                        Sync
                    </button>
                </div>
            </div>
        </div>

        {#if games.length == 0}
            <div></div>
        {:else}
            <div class="container mx-auto">
                <div class="flex flex-wrap -mx-2">
                    {#each games as game}
                        <div
                            class="w-full sm:w-1/2 md:w-1/3 lg:w-1/4 px-2 pt-3"
                        >
                            <div
                                class="bg-opacity-50 rounded-lg overflow-hidden shadow-lg"
                            >
                                <img
                                    src={game.cover_hd_url}
                                    alt=""
                                    class="rounded-b-lg w-full"
                                />
                                <div class="text-white pt-2 text-center">
                                    {game.name}
                                </div>
                            </div>
                        </div>
                    {/each}
                </div>
            </div>
        {/if}

        {#if fetchinGames}
            <div class="text-white text-2xl text-center">LOADING...</div>
        {:else}
            <div></div>
        {/if}
    </div>
</main>
