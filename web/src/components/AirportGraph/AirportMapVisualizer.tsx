import React from "react"
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { AirportMap } from "@/api"

type AirportMapProps = {
  airportMap: AirportMap
}

export function AirportMapVisualizer({ airportMap }: AirportMapProps) {
  const { nodes, edges } = airportMap

  // --- 1) Compute a circular layout for nodes ---
  const centerX = 500
  const centerY = 500
  const radius = Math.max(200, 15 * nodes.length) // Make the circle bigger for many nodes

  // Generate (x, y) positions in a circle for each node
  const nodePositions = new Map<string, { x: number; y: number }>()
  nodes.forEach((node, index) => {
    const angle = (2 * Math.PI * index) / nodes.length
    const x = centerX + radius * Math.cos(angle)
    const y = centerY + radius * Math.sin(angle)
    nodePositions.set(node.id, { x, y })
  })

  return (
    <Card>
      <CardHeader>
        <CardTitle>Airport Map</CardTitle>
      </CardHeader>
      <CardContent>
        {/* 
          You can style or size the SVG however you like.
          For example, set a large width/height or 
          make it responsive via a wrapper.
        */}
        <svg width={1000} height={1000} style={{ border: "1px solid #ccc" }}>
          {/* --- 2) Draw edges as lines --- */}
          {edges.map((edge, idx) => {
            const fromPos = nodePositions.get(edge.from)
            const toPos = nodePositions.get(edge.to)
            if (!fromPos || !toPos) return null

            return (
              <line
                key={idx}
                x1={fromPos.x}
                y1={fromPos.y}
                x2={toPos.x}
                y2={toPos.y}
                stroke="#555"
                strokeWidth={2}
              />
            )
          })}

          {/* --- 3) Draw nodes as circles + labels --- */}
          {nodes.map((node) => {
            const pos = nodePositions.get(node.id)!
            return (
              <g key={node.id}>
                <circle
                  cx={pos.x}
                  cy={pos.y}
                  r={10}
                  fill="#007bff"
                  stroke="#000"
                  strokeWidth={1}
                />
                {/* Node ID label */}
                <text
                  x={pos.x}
                  y={pos.y - 15}
                  textAnchor="middle"
                  fontSize={12}
                  fill="#333"
                >
                  {node.id}
                </text>
              </g>
            )
          })}
        </svg>
      </CardContent>
    </Card>
  )
}
