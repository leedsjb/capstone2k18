import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";

import FlexFillVH from "../../components/FlexFillVH";
import LoadingSpinner from "../../components/LoadingSpinner";
import PeopleDetailsItem from "../../components/PeopleDetailsItem";
import ScrollView from "../../components/ScrollView";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";
import Error from "../../components/Error";

import { fetchPeopleDetail } from "../../actions/peopleDetail/actions";

class PeopleDetailPage extends Component {
    componentDidMount() {
        if (this.props.id) {
            this.props.fetchPeopleDetail(this.props.id);
        }
    }

    renderPeopleDetail() {
        if (
            !this.props.peopleDetail.pending &&
            !Array.isArray(this.props.peopleDetail.data)
        ) {
            return <PeopleDetailsItem person={this.props.peopleDetail.data} />;
        }

        return <LoadingSpinner />;
    }

    render() {
        let backPath =
            new URLSearchParams(window.location.search).get("source") ===
            "groups"
                ? `/groups/${new URLSearchParams(window.location.search).get(
                      "id"
                  )}`
                : "/people";
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>People</title>
                </Helmet>
                <TitleBar back backPath={backPath} title="People" />
                {this.props.peopleDetail.error ? (
                    <Flex
                        flexDirection="column"
                        flex={1}
                        alignItems="center"
                        justifyContent="center"
                    >
                        <Error
                            title="An error has occurred"
                            content={this.props.peopleDetail.error.toString()}
                        />
                    </Flex>
                ) : (
                    <ScrollView>{this.renderPeopleDetail()}</ScrollView>
                )}

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
