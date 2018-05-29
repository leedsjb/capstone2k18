import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import NavBar from "../../components/NavBar";
import FlexFillVH from "../../components/FlexFillVH";
import Error from "../../components/Error";

class NotFoundPage extends Component {
    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Resources</title>
                </Helmet>

                <TitleBar title="Resources" />
                <NavBar />

                <Flex
                    flexDirection="column"
                    flex={1}
                    alignItems="center"
                    justifyContent="center"
                >
                    <Error notFound />
                </Flex>

                <TabBar />
            </FlexFillVH>
        );
    }
}

export default NotFoundPage;
