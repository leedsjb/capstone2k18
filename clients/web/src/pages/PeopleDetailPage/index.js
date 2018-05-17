import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";

import Box from "../../components/Box";
import ButtonIcon from "../../components/ButtonIcon";
import FlexFillVH from "../../components/FlexFillVH";
import Heading from "../../components/Heading";
import ProfileAvatar from "../../components/ProfileAvatar";
import ScrollView from "../../components/ScrollView";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";

import { fetchPeopleDetail } from "../../actions/peopleDetail/actions";

class PeopleDetailPage extends Component {
    componentDidMount() {
        if (this.props.id) {
            this.props.fetchPeopleDetail(this.props.id);
        }
    }

    renderPeopleDetail(person) {
        if (
            !this.props.peopleDetail.pending &&
            !Array.isArray(this.props.peopleDetail.data)
        ) {
            return (
                <Flex flexDirection="column" alignItems="center">
                    <Box mt={4}>
                        <ProfileAvatar fName={person.fName} size={72} />
                    </Box>
                    <Heading
                        children={`${person.fName} ${person.lName}`}
                        is="h2"
                        fontSize={4}
                        mt={3}
                    />
                    <Heading
                        children={`${person.position}`}
                        is="h3"
                        fontWeight="normal"
                        fontSize={2}
                    />
                    <Flex mt={3}>
                        <ButtonIcon glyph="bubbleChat">Text</ButtonIcon>
                        <Box mx={3}>
                            <ButtonIcon glyph="phone">Call</ButtonIcon>
                        </Box>
                        <ButtonIcon glyph="email">Mail</ButtonIcon>
                    </Flex>
                </Flex>
            );
        }

        return <div>Loading...</div>;
    }

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>People</title>
                </Helmet>
                <TitleBar back backPath={"/people"} title="People" />
                <ScrollView>
                    {this.renderPeopleDetail(this.props.peopleDetail.data)}
                </ScrollView>
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        peopleDetail: state.peopleDetail,
        id: ownProps.match.params.id
    };
}

const mapDispatchToProps = {
    fetchPeopleDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(PeopleDetailPage);
