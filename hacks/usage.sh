#!/bin/bash

export AUTH="eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlFtSXdJNVprRjBJTmh0b0dwVWFmcyJ9.eyJuaWNrbmFtZSI6ImZlbGl4IiwibmFtZSI6ImZlbGl4QHBvanRpbmdlci5jb20iLCJwaWN0dXJlIjoiaHR0cHM6Ly9zLmdyYXZhdGFyLmNvbS9hdmF0YXIvZGI4NTZkZjMzZmE0ZjRiY2U0NDE4MTlmNjA0YzkwZDU_cz00ODAmcj1wZyZkPWh0dHBzJTNBJTJGJTJGY2RuLmF1dGgwLmNvbSUyRmF2YXRhcnMlMkZmZS5wbmciLCJ1cGRhdGVkX2F0IjoiMjAyMS0wMy0wNlQxNjowMzoyOS43MjZaIiwiZW1haWwiOiJmZWxpeEBwb2p0aW5nZXIuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImlzcyI6Imh0dHBzOi8vcG9qbnRmeC5ldS5hdXRoMC5jb20vIiwic3ViIjoiYXV0aDB8NWY3Zjc0YmY4ODU0N2YwMDZmMGRlMWJhIiwiYXVkIjoicTMxSVVwQWc1UXUxOGVCTWlCc2JJdkg5aGhScTdXZ0UiLCJpYXQiOjE2MTUxNDcyOTcsImV4cCI6MTYxNTE4MzI5N30.W7ccbmOh81vuUo8KNNwfs8zVDOLl4_nJG9brqn9r1HbhozwBbKttANYlq3FcJPqyE5zgiy24hA6FGAkR7xy4XJzLTaIqWDIusvL36O_PAarLRFXUVzCZVS92D6uemVybqNW3BVZXh0Euub8Y9x4fCkSpZoHYQ34Ad1Kr4BSMMVL3yzOQly7Te--2mC_RP3320KBMXc93S8DKJ4DCE9URejom7nc_dgnQa0zgYzs54q7k-Lho-gsPa2AAfERMHkJLiN3xqeaG5p6Ag1PqALg5Ssyw0QSYWmVu-4-I2SA_oGs7_6GHzLNtmr10gwqFQzag2_Zm7E-N0IYeT1BHcNc_Zw"

grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "NodeScanTimeout": 500, "PortScanTimeout": 50, "MACAddress": "" }' -plaintext localhost:15123 com.pojtinger.felicitas.liwasc.NodeAndPortScanService.StartNodeScan
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "NodeScanTimeout": 500, "PortScanTimeout": 50, "MACAddress": "02:42:64:96:47:8a" }' -plaintext localhost:15123 com.pojtinger.felicitas.liwasc.NodeAndPortScanService.StartNodeScan
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext localhost:15123 com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToNodeScans
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{"ID": 11}' localhost:15123 com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToNodes
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{"ID": 11}' localhost:15123 com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToPortScans
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{"ID": 11}' localhost:15123 com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToPorts
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext localhost:15123 com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForScanner
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "MACAddress": "02:42:64:96:47:8a" }' localhost:15123 com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForNode
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "PortNumber": 7005, "TransportProtocol": "tcp" }' localhost:15123 com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForPort
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "NodeWakeTimeout": 1000, "MACAddress": "02:42:64:96:47:8a" }' localhost:15123 com.pojtinger.felicitas.liwasc.NodeWakeService.StartNodeWake
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext localhost:15123 com.pojtinger.felicitas.liwasc.NodeWakeService.SubscribeToNodeWakes
