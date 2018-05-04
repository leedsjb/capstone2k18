import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchPeopleDetail } from "../../actions/peopleDetail/actions";

class PeopleDetailReducer extends Component {
    componentDidMount() {
        this.props.fetchPeopleDetail(this.props.id);
    }

    render() {
        return this.props.render({ peopleDetail: this.props.peopleDetail });
    }
}

const mapStateToProps = ({ peopleDetail }) => {
    return {
        peopleDetail
    };
};

const mapDispatchToProps = {
    fetchPeopleDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(
    PeopleDetailReducer
);
