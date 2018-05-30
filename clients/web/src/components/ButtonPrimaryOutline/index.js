import sys from "system-components";
import ButtonPrimaryGradient from "../ButtonPrimaryGradient";

const ButtonPrimaryOutline = sys(
    {
        is: ButtonPrimaryGradient,
        py: 3,
        px: 6,
        background: "none",
        color: "black1",
        border: "1px solid transparent"
    },
    "space",
    props => ({
        border: `1px solid ${props.theme.colors.airlift1}`,
        "&:hover": {
            background: `radial-gradient(circle at top left, ${
                props.theme.colors.airlift1
            }, ${props.theme.colors.airlift2})`,
            color: "white"
        }
    })
);

ButtonPrimaryOutline.displayName = "ButtonPrimaryOutline";

export default ButtonPrimaryOutline;
