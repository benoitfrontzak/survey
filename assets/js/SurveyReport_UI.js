class SurveyReport_UI {
    // drawChart(data, options, targetID) {

    //     var view = new google.visualization.DataView(data);
    //     view.setColumns([0, 1,
    //                      { calc: "stringify",
    //                        sourceColumn: 1,
    //                        type: "string",
    //                        role: "annotation" },
    //                      2]);
  
    //     var chart = new google.visualization.ColumnChart(document.getElementById(targetID));
    //     chart.draw(view, options);
    // }
    drawChart(myData, title, targetID) {
        var data = google.visualization.arrayToDataTable(myData);
  
        var view = new google.visualization.DataView(data);
        view.setColumns([0, 1,
                         { calc: "stringify",
                           sourceColumn: 1,
                           type: "string",
                           role: "annotation" },
                         2]);
        var options = {
            title: title,
            width: 600,
            height: 400,
            bar: {groupWidth: "95%"},
            legend: { position: "none" },
        }

        var chart = new google.visualization.ColumnChart(document.getElementById(targetID));
        chart.draw(view, options);
    }
}
