import React from "react";
import Media from "react-media";
import { withTheme } from "styled-components";

const DetailView = ({ children, theme: { breakpoints } }) => {
    return (
        <Media
            query={`(min-width: ${breakpoints[1]})`}
            render={() => {
                return children;
            }}
        />
    );
};

export default withTheme(DetailView);
