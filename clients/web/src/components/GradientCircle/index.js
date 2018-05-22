import sys from "system-components";
import Circle from "../Circle";

const GradientCircle = sys(
    {
        is: Circle
    },
    props => ({
        background: `radial-gradient(circle at top left, ${props.firstColor}, ${
            props.secondColor
        })`
    })
);

export default GradientCircle;
