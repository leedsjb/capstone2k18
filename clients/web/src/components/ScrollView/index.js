import styled from "styled-components";
import { width, maxWidth, borders } from "styled-system";

const ScrollView = styled.div`
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow-y: auto;
    height: 100%;
    ${width};
    ${maxWidth};
    ${borders};
`;

export default ScrollView;
