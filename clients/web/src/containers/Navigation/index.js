import React, { Component } from "react";
import { Box } from "grid-styled";
import { Link as RouterLink } from "react-router-dom";

import Link from "../../components/Link";

class Navigation extends Component {
    render() {
        return (
            <Box bg="wireframe" py={3}>
                <Link is={RouterLink} to="/">
                    AirliftNW
                </Link>
                <Link is={RouterLink} to="/missions">
                    Missions
                </Link>
                <Link is={RouterLink} to="/personnel">
                    Personnel
                </Link>
            </Box>
        );
    }
}

export default Navigation;
