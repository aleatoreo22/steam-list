import type { Game } from '../model/game'

const url = 'http://127.0.0.1:8080/api/';

async function getPlayerGames(playerId: string, page: number = 1): Promise<Game[]> {
    let endpoint = `${url}player/games/${playerId}?page=${page}`
    console.log(endpoint)
    let json = await fetch(endpoint)
        .then(response => response.json())
        .then(data => data)
        .catch(error => console.error(error));
    let games = json as Game[]
    console.log(json)
    return games
}

export { getPlayerGames }