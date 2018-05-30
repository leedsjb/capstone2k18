import sys from "system-components";
import Button from "../Button";
import { themeGet } from "styled-system";

const ButtonPrimary = sys(
    {
        is: Button,
        py: 3,
        px: 4
    },
    "space",
    props => ({
        backgroundColor: `${themeGet("colors.black2")(props)}`,
        color: "white"
    })
);

ButtonPrimary.displayName = "ButtonPrimary";

export default ButtonPrimary;
