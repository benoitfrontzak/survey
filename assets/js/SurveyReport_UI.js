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

    CreatePDFfromHTML() {
        var HTML_Width = $(".html-content").width();
        var HTML_Height = $(".html-content").height();
        var top_left_margin = 15;
        var PDF_Width = HTML_Width + (top_left_margin * 2);
        var PDF_Height = (PDF_Width * 1.5) + (top_left_margin * 2);
        var canvas_image_width = HTML_Width;
        var canvas_image_height = HTML_Height;
    
        var totalPDFPages = Math.ceil(HTML_Height / PDF_Height) - 1;
    
        html2canvas($(".html-content")[0]).then(function (canvas) {
            var imgData = canvas.toDataURL("image/jpeg", 1.0);
            var pdf = new jsPDF('p', 'pt', [PDF_Width, PDF_Height]);
            pdf.addImage(imgData, 'JPG', top_left_margin, top_left_margin, canvas_image_width, canvas_image_height);
            for (var i = 1; i <= totalPDFPages; i++) { 
                pdf.addPage(PDF_Width, PDF_Height);
                pdf.addImage(imgData, 'JPG', top_left_margin, -(PDF_Height*i)+(top_left_margin*4),canvas_image_width,canvas_image_height);
            }
            pdf.save("Your_PDF_Name.pdf");
            $(".html-content").hide();
        });
    }
}
