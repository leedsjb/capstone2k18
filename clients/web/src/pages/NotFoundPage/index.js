import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

import TitleBar from "../../components/TitleBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import NavBar from "../../components/NavBar";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";
import Box from "../../components/Box";
import Image from "../../components/Image";
import Absolute from "../../components/Absolute";
import FlexFillHeight from "../../components/FlexFillHeight";
import Relative from "../../components/Relative";
import Span from "../../components/Span";
import Card from "../../components/Card";
import Error from "../../components/Error";
import LoadingSpinner from "../../components/LoadingSpinner";

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
