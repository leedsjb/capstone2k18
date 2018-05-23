// Adaptation of https://raw.githubusercontent.com/jxnblk/rebass/master/src/ButtonOutline.js

import styled from "styled-components";
import { themeGet } from "styled-system";
import Button from "../Button";

const ButtonOutline = styled(Button)([], props => ({
    background: "#F7F9FC",
    boxShadow: `inset 0 0 0 2px`,
    "&:hover": {
        color: "#C20B29",
        backgroundColor: themeGet("colors.blue")(props)
    },
    "&:focus": {
        boxShadow: `inset 0 0 0 2px, 0 0 0 2px`
    },
    "&:active": {
        color: "white",
        backgroundColor: themeGet("colors.blue")(props),
        boxShadow: `inset 0 0 0 2px ${themeGet("colors." + props.color)(
            props
        )}, inset 0 0 8px ${themeGet("colors.darken.1")(props)}`
    }
}));

ButtonOutline.displayName = "ButtonOutline";

ButtonOutline.defaultProps = {
    color: "#dc0c2f",
    bg: "transparent"
};

export default ButtonOutline;
