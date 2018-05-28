import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import { Helmet } from "react-helmet";

import Box from "../../components/Box";
import Divider from "../../components/Divider";
import Heading from "../../components/Heading";
import FlexFillVH from "../../components/FlexFillVH";
import PeopleListItem from "../../components/PeopleListItem";
import PeopleLoader from "../../components/PeopleLoader";
import ScrollView from "../../components/ScrollView";
import Span from "../../components/Span";
import TabBar from "../../components/TabBar";
import Text from "../../components/Text";
import TitleBar from "../../components/TitleBar";

import { fetchGroupsDetail } from "../../actions/groupsDetail/actions";

class GroupsDetailPage extends Component {
    componentDidMount() {
        if (this.props.id) {
            this.props.fetchGroupsDetail(this.props.id);
        }
    }

    renderGroupsDetail() {
        if (
            !this.props.groupsDetail.pending &&
            !Array.isArray(this.props.groupsDetail.data)
        ) {
            return this.props.groupsDetail.data.people.map((person, i) => {
                return (
                    <div key={person.id}>
                        <Link
                            to={`/people/${person.id}?source=groups&id=${
                                this.props.id
                            }`}
                        >
                            <PeopleListItem
                                active={
                                    Number(this.props.id) === person.id ? 1 : 0
                                }
                                person={person}
                            />
                        </Link>
                        {this.props.groupsDetail.data.length === 1 ||
                        i !== this.props.groupsDetail.data.length - 1 ? (
                            <Divider />
                        ) : null}
                    </div>
                );
            });
        } else if (!this.props.groupsDetail.pending) {
            return (
                <Box mt={4}>
                    <Heading is="h2" fontSize={4} textAlign="center">
                        No Group Details
                    </Heading>
                    <Text textAlign="center">Empty State Text</Text>
                </Box>
            );
        } else if (this.props.groupsDetail.pending) {
            return (
                <Box mt={3}>
                    <PeopleLoader />
                    <PeopleLoader />
                    <PeopleLoader />
                </Box>
            );
        }
    }

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Groups</title>
                </Helmet>
                <TitleBar back backPath="/groups" title="Groups" />
                {this.props.groupsDetail.error ? (
                    <FlexFillVH>
                        An error has occurred:{" "}
                        {this.props.groupsDetail.error.toString()}
                    </FlexFillVH>
                ) : (
                    <FlexFillVH flexDirection="column">
                        <Flex justifyContent="center" py={2} px={3}>
                            <Span fontWeight="bold">
                                {this.props.groupsDetail.data.name}
                            </Span>
                        </Flex>
                        <Divider />
                        <ScrollView>{this.renderGroupsDetail()}</ScrollView>
                    </FlexFillVH>
                )}
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        groupsDetail: state.groupsDetail,
        id: ownProps.match.params.groupID
    };
}

const mapDispatchToProps = {
    fetchGroupsDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(GroupsDetailPage);
