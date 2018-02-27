import React, { Component } from "react";
import { Link } from "react-router-dom";

class Navigation extends Component {
    render() {
        return (
            <div>
                <Link to="/">AirliftNW</Link>
                <Link to="/missions">Missions</Link>
                <Link to="/personnel">Personnel</Link>
            </div>
        );
    }
}

export default Navigation;
