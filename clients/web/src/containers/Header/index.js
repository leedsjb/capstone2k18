import React, { Component } from "react";
import { Link } from "react-router-dom";

class Header extends Component {
    render() {
        return (
            <div>
                <Link to="/">AirliftNW</Link>
            </div>
        );
    }
}

export default Header;
