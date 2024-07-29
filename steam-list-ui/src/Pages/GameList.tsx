import { Cover, CoverProps } from "../Components/Cover";

export function GameList() {
  const coverProps: CoverProps = {
    cover_hd_url: "//images.igdb.com/igdb/image/upload/t_cover_big/co8fpz.jpg",
    name: "Oblivion",
  };

  return <Cover {...coverProps} />;
}
