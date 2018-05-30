import sys from "system-components";
import Button from "../Button";

const ButtonPrimaryGradient = sys(
    {
        is: Button,
        py: 3,
        px: 6,
        border: "1px solid transparent"
    },
    "space",
    props => ({
        background: `radial-gradient(circle at top left, ${
            props.theme.colors.airlift1
        }, ${props.theme.colors.airlift2})`
    })
);

ButtonPrimaryGradient.displayName = "ButtonPrimaryGradient";

export default ButtonPrimaryGradient;
