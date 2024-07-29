import { Box, Paper } from "@mui/material";
import { Game } from "../models/game";

export interface CoverProps extends Game {}

export function Cover(props: CoverProps) {
  return (
    <Paper
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        margin: "20px",
        paddingTop: "8px",
        borderRadius: "20px",
        bgcolor: "#171717",
        paddingX: "8px",
      }}
    >
      <img
        style={{ borderRadius: "20px" }}
        src={props.cover_hd_url}
        height={"200px"}
      />
      <h3 style={{ color: "white" }}>{props.name}</h3>
    </Paper>
  );
}
