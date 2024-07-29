import { Android } from "@mui/icons-material";
import {
  AppBar,
  Avatar,
  Box,
  IconButton,
  Toolbar,
  Tooltip,
  Typography,
} from "@mui/material";

export function HeaderBar() {
  return (
    <AppBar sx={{ bgcolor: "#1c1c1c" }}>
      <Box sx={{ width: "100vw" }}>
        <Toolbar>
          <Android
            sx={{
              display: { xs: "none", md: "flex" },
              mr: 2,
            }}
          />
          <Typography
            variant="h5"
            sx={{
              display: { xs: "none", md: "flex" },
              flexGrow: 1,
              fontWeight: 700,
              letterSpacing: ".1rem",
              color: "inherit",
            }}
          >
            My Steam List
          </Typography>
          <Box sx={{ justifyContent: "flex-end", display: "inline-flex" }}>
            <Tooltip title="Open settings">
              <IconButton sx={{ p: 0 }}>
                <Avatar alt="Remy Sharp" src="/static/images/avatar/2.jpg" />
              </IconButton>
            </Tooltip>
          </Box>
        </Toolbar>
      </Box>
    </AppBar>
  );
}
