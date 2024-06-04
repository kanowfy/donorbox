const formatNumber = (num) => {
    return num.toLocaleString();
}

const calculateProgress = (current, goal) => {
    const progress = Math.floor(current / goal * 100);
    return progress < 100 ? progress : 100;
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

const formatDateTime = (date) => {
    let year = date.getFullYear();
    let month = (1 + date.getMonth()).toString().padStart(2, '0');
    let day = date.getDate().toString().padStart(2, '0');
    let hour = date.getHours().toString().padStart(2, '0');
    let minute = date.getMinutes().toString().padStart(2, '0');
    let second = date.getSeconds().toString().padStart(2, '0');

    return day + '/' + month + '/' + year + " " + hour + ":" + minute + ":" + second;
}

const parseExpiry = (date) => {
    return {
        month: parseInt(date.substring(0, 2), 10),
        year: parseInt("20" + date.substring(2), 10)
    }
}

export default { formatNumber, calculateProgress, calculateDayDifference, parseDateFromRFC3339, getDaySince, formatDate, formatDateTime, parseExpiry }
