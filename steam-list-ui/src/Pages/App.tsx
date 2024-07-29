import { Box, Toolbar } from "@mui/material";
import { HeaderBar } from "../Components/HeaderBar";
import { ReactNode } from "react";

export interface AppProps {
  children: ReactNode;
}

export function App(props: AppProps) {
  return (
    <Box>
      <HeaderBar />
      <Toolbar />
      {props.children}
    </Box>
  );
}
