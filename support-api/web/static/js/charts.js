document.addEventListener('DOMContentLoaded', function() {
    updateDashboardTitle();
    initTicketServicesChart();
    initTicketStatesChart();
    initArticleSendersChart();
    initArticleTypesChart();
    initArticleTimesChart();
});

function getUrlParams() {
    const params = new URLSearchParams(window.location.search);
    const defaultFrom = '2025-02-01'; 
    const defaultTo = '2025-02-28';
    
    return {
        from: params.get('from') || defaultFrom,
        to: params.get('to') || defaultTo
    };
}

function updateDashboardTitle() {
    const {from, to} = getUrlParams();
    const titleElement = document.querySelector('h1');
    
    if (titleElement) {
        const originalText = titleElement.textContent;
        titleElement.textContent = `${originalText} ${from} - ${to}`;
    }
}
function initTicketServicesChart() {
    const {from, to} = getUrlParams();

    fetch(`/api/tickets/services?from=${from}&to=${to}`)
        .then(response => response.json())
        .then(data => {
            new Chart(
                document.getElementById('ticketServicesChart'),
                {
                    type: 'bar',
                    data: data,
                    options: {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'top',
                            },
                            title: {
                                display: true,
                                text: 'Ticket Services Distribution'
                            }
                        },
                    }
                }
            );
        });
}


function initTicketStatesChart() {
    const {from, to} = getUrlParams();

    fetch(`/api/tickets/states?from=${from}&to=${to}`)
        .then(response => response.json())
        .then(data => {
            new Chart(
                document.getElementById('ticketStatesChart'),
                {
                    type: 'bar',
                    data: data,
                    options: {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'top',
                            },
                            title: {
                                display: true,
                                text: 'Ticket States Distribution'
                            }
                        },
                    }
                }
            );
        });
}

function initArticleTypesChart() {
    const {from, to} = getUrlParams();


    fetch(`/api/articles/types?from=${from}&to=${to}`)
        .then(response => response.json())
        .then(data => {
            new Chart(
                document.getElementById('articleTypesChart'),
                {
                    type: 'pie',
                    data: data,
                    options: {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'top',
                            },
                            title: {
                                display: true,
                                text: 'Article Types Distribution'
                            }
                        }
                    }
                }
            );
        });
}

function initArticleSendersChart() {
    const {from, to} = getUrlParams();


    fetch(`/api/articles/senders?from=${from}&to=${to}`)
        .then(response => response.json())
        .then(data => {
            new Chart(
                document.getElementById('articleSendersChart'),
                {
                    type: 'pie',
                    data: data,
                    options: {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'top',
                            },
                            title: {
                                display: true,
                                text: 'Article Senders Distribution'
                            }
                        }
                    }
                }
            );
        });
}

function initArticleTimesChart() {
    const {from, to} = getUrlParams();

    fetch(`/api/articles/create-time?from=${from}&to=${to}`)
        .then(response => response.json())
        .then(data => {
            const ctx = document.getElementById('articleTimesChart').getContext('2d');
            
            new Chart(ctx, {
                type: 'line',
                data: {
                    labels: data.labels,
                    datasets: [{
                        label: data.datasets[0].label,
                        data: data.datasets[0].data,
                        borderColor: '#36b9cc',
                        backgroundColor: 'rgba(54, 185, 204, 0.1)',
                        borderWidth: 2,
                        fill: true
                    }]
                },
                options: {
                    responsive: true,
                    scales: {
                        x: {
                            type: 'time',
                            time: {
                                parser: 'yyyy-MM-dd\'T\'HH:mm:ss\'Z\'',
                                tooltipFormat: 'dd MMM yyyy',
                                unit: 'day',
                                displayFormats: {
                                    day: 'dd MMM yyyy'
                                }
                            },
                            ticks: {
                                autoSkip: true,
                                maxTicksLimit: 10
                            }
                        },
                        y: {
                            beginAtZero: true
                        }
                    }
                }
            });
        });
}