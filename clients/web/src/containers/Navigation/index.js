import React, { Component } from "react";
import { Link } from "react-router-dom";
import { Box } from "grid-styled";

class Navigation extends Component {
    render() {
        return (
            <Box bg="wireframe" py={3}>
                <Link to="/">AirliftNW</Link>
                <Link to="/missions">Missions</Link>
                <Link to="/personnel">Personnel</Link>
            </Box>
        );
    }
}

export default Navigation;
