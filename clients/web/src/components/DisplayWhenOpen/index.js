import styled from "styled-components";
import Box from "../Box";

const Displayisopen = styled(Box)`
    display: ${props => (props.isopen ? "block" : "none")};
`;

export default Displayisopen;
