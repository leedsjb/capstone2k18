import styled from "styled-components";

import Box from "../Box";

const ScrollView = styled(Box)`
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow-y: auto;
    height: 100%;
`;

export default ScrollView;
