# **D3.js** 
**D3.js (Data-Driven Documents)** is a JavaScript library that allows you to bring data to life
using web standards. It provides a flexible structure for manipulating documents
based on data and is known for its ability to aid in the creation of sophisticated and custom data visualizations in the browser.

***

## **Key Features** :
- **Data Binding** : D3 allows you to bind arbitrary data to a Document Object-Model (DOM), and then apply data-driven
transformations to the document.
- **Powerful Visualization Components** : Whether it's a bar chart , pie chart, or a geographical map, D3
got you covered.
- **Custom Visualization** : Beyond the basics, D3 provides the flexibility to design
unique visualizations.
- **Dynamic Properties** : D3 Lets you apply transformations based on functions. For example
you can transform the radius of a circle based on data.

## **Installation** :
You can include D3.js in your project using multiple methods:

1. Using a Content Delivery Network (CDN) :
```
<script src="https://d3js.org/d3.v7.min.js"></script>
```
2. Using npm (Node Package Manager) :
```
npm install d3
```

## Basic Usage :
1. **Binding data** :
   D3 allows you to bind data to elements. For instance, you can bind
   an array of numbers to paragraphs :
   ##### JavaScript
   ```
   d3.select("body")
   .selectAll("p")
   .data([4,8,16,23,42])
   .enter().append("p")
   .text(function(d) {return "This is number " + d: });
   ```
   
2. **Creating a Simple Bar Chart** :
   Assuming you have a block of SVG in your HTML :
   #### html
   ```
   <svg width="500" height="500"></svg>
   ```
   You can create a bar chart with D3 as follows :
   #### JavaScript
   ```
   const data = [30, 86, 168m 281, 303, 365];

   const svg = d3.select("svg"),
               width = svg.sttr("width"),
               scaleFactor = 10,
               barHeight = 20;
   
   const bar = svg.selectAll("g")
         .data(data)
         .enter().append("g")
         .attr("transform", (d, i) => "translate(0," + i * barHeight + ")");

   bar.append("rect")
       .attr("width", d => d * scaleFactor)
       .attr("height", barHeight - 1);

   bar.append("text")
       .attr("x",  d => (d * scaleFactor) - 3)
       .attr("y", barHeight / 2)
       .attr("dy", ".35cm")
       .text(d => d);
   ```

3. **Interactivity** :
   Once of the strenghts of D3 is the ability to add interactivity :
   #### JavaScript
   ```
   d3.select("body")
     .selectAll("p")
     .data([4,8,15,16,23])
     .enter().append("p")
     .text(function(d) { return "This is number" + d; })
     .on("click", function(d) {console.log("You clicked on number" + d});
   ```

## Project Example
- [Complex bar chart](https://github.com/rifqanzalbina/DifferentButSame/tree/main/Complexd3chart)

## **Conclusion**
<p>
This overview only scratches the surface of D3's capabilities. For more intricate and detailed visuals, the D3 community and official documentation offer countless examples, tutorials, and resources. The flexibility of D3 is both its strength and its challenge, as the possibilities are vast but require a deeper understanding to fully exploit. If you're interested in data visualization in the web domain, 
spending time learning D3 is a worthy investment.
</p>
   
    
   
