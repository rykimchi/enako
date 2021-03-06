import React, { Component } from 'react';
import DetailList from './DetailList';
import Card from './Card';
import moment from 'moment';
import '../css/Details.css';

class Details extends Component {
    filterExpenses = (date) => {
        const compareDate = moment(date).format('YYYY-MM-DD');
        const expenses = this.props.expenses.filter(expense => expense.expense_date === compareDate);
        if (expenses.length) {
            return expenses;
        }

        return [];
    };

    render() {
        return (
            <div className='Details'>
                <Card heading={moment(this.props.selectedDate).format('MMMM Do YYYY')}>
                    <DetailList
                        selectedDate={this.props.selectedDate}
                        types={this.props.types}
                        categories={this.props.categories}
                        expenses={this.filterExpenses(this.props.selectedDate)}
                    />
                </Card>
            </div>
        );
    }
}

export default Details;
