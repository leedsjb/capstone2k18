// Separate component as we might need to change it based on screen width

import styled from "styled-components";

const ScrollView = styled.div`
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow-y: auto;
`;

export default ScrollView;
