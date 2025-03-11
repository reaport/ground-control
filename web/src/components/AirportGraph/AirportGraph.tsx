import React, { useEffect, useState, useRef } from "react";
import ForceGraph, { GraphData, NodeObject, LinkObject } from "react-force-graph-2d";
import { AirportMap } from "../../api";

interface AirportGraphProps {
  data: AirportMap;
}

export default function AirportGraph({ data }: AirportGraphProps) {
  const [graphData, setGraphData] = useState<GraphData>({ nodes: [], links: [] });
  const [hoverNode, setHoverNode] = useState<NodeObject | null>(null);
  const hoverNodeRef = useRef<NodeObject | null>(null);

  useEffect(() => {
    if (data) {
      const nodes: NodeObject[] = data.nodes.map((node) => ({ 
        id: node.id, 
        group: node.types.length, 
        vehicles: node.vehicles,
        hasVehicles: node.vehicles && node.vehicles.length > 0
      }));
      const links: LinkObject[] = data.edges.map((edge) => ({ source: edge.from, target: edge.to, distance: edge.distance }));
      setGraphData({ nodes, links });
    }
  }, [data]);

  return (
    <div style={{ width: "100vw", height: "100vh", position: "relative" }}>
      {hoverNode && hoverNode.vehicles && hoverNode.vehicles.length > 0 && (
        <div style={{
          position: "absolute",
          top: 10,
          left: 10,
          backgroundColor: "rgba(0, 0, 0, 0.75)",
          color: "white",
          padding: "8px",
          borderRadius: "4px",
        }}>
          <strong>Vehicles at {hoverNode.id}:</strong>
          <ul>
            {hoverNode.vehicles.map((v) => (
              <li key={v.id}>{v.id} ({v.type})</li>
            ))}
          </ul>
        </div>
      )}
      <ForceGraph
        graphData={graphData}
        nodeAutoColorBy={(node: any) => node.hasVehicles ? "red" : "group"}
        linkDirectionalArrowLength={5}
        linkDirectionalArrowRelPos={1}
        nodeCanvasObject={(node: any, ctx: CanvasRenderingContext2D, globalScale: number) => {
          const label = node.id;
          const fontSize = 12 / globalScale;
          ctx.font = `${fontSize}px Sans-Serif`;
          ctx.fillStyle = "black";
          ctx.fillText(label, node.x + 6, node.y + 6);
          
          ctx.fillStyle = node.hasVehicles ? "red" : "blue";
          ctx.beginPath();
          ctx.arc(node.x, node.y, 5, 0, 2 * Math.PI, false);
          ctx.fill();
        }}
        onNodeHover={(node: any) => {
          hoverNodeRef.current = node || null;
          setHoverNode(node || null);
        }}
        linkLabel={(link: LinkObject) => `Distance: ${link.distance}`}
      />
    </div>
  );
}