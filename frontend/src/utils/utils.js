const formatNumber = (num) => {
    return num.toLocaleString();
}

const calculateProgress = (current, goal) => {
    return Math.floor(current / goal * 100);
}

const calculateDayDifference = (dateA, dateB) => {
    return Math.round((dateA - dateB) / (1000 * 60 * 60 * 24));
}

const parseDateFromRFC3339 = (date) => {
    return Date.parse(date)
}

const getDaySince = (date) => {
    return calculateDayDifference(Date.now(), parseDateFromRFC3339(date));
}

const formatDate = (date) => {
    let year = date.getFullYear();
    let month = (1 + date.getMonth()).toString().padStart(2, '0');
    let day = date.getDate().toString().padStart(2, '0');

    return day + '/' + month + '/' + year;
}

export default { formatNumber, calculateProgress, calculateDayDifference, parseDateFromRFC3339, getDaySince, formatDate }
