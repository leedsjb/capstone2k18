import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchProfile } from "../../actions/profile/actions";

class ProfileProvider extends Component {
    componentDidMount() {
        this.props.fetchProfile();
    }

    render() {
        return this.props.render({ profile: this.props.profile });
    }
}

const mapStateToProps = ({ profile }) => {
    return {
        profile
    };
};

const mapDispatchToProps = {
    fetchProfile
};

export default connect(mapStateToProps, mapDispatchToProps)(ProfileProvider);
