import styled from "styled-components";
import { width, maxWidth } from "styled-system";

const ScrollView = styled.div`
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow-y: auto;
    ${width};
    ${maxWidth};
`;

export default ScrollView;
