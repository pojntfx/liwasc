#!/bin/bash

export AUTH='eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlFtSXdJNVprRjBJTmh0b0dwVWFmcyJ9.eyJuaWNrbmFtZSI6ImZlbGl4IiwibmFtZSI6ImZlbGl4QHBvanRpbmdlci5jb20iLCJwaWN0dXJlIjoiaHR0cHM6Ly9zLmdyYXZhdGFyLmNvbS9hdmF0YXIvZGI4NTZkZjMzZmE0ZjRiY2U0NDE4MTlmNjA0YzkwZDU_cz00ODAmcj1wZyZkPWh0dHBzJTNBJTJGJTJGY2RuLmF1dGgwLmNvbSUyRmF2YXRhcnMlMkZmZS5wbmciLCJ1cGRhdGVkX2F0IjoiMjAyMS0wMy0wNlQxNjowMzoyOS43MjZaIiwiZW1haWwiOiJmZWxpeEBwb2p0aW5nZXIuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImlzcyI6Imh0dHBzOi8vcG9qbnRmeC5ldS5hdXRoMC5jb20vIiwic3ViIjoiYXV0aDB8NWY3Zjc0YmY4ODU0N2YwMDZmMGRlMWJhIiwiYXVkIjoicTMxSVVwQWc1UXUxOGVCTWlCc2JJdkg5aGhScTdXZ0UiLCJpYXQiOjE2MTUwNDY2MTEsImV4cCI6MTYxNTA4MjYxMX0.tf-mYEhXYlMmJy4tDfheYpnZ4O6hh-C3alWCOcdnJtA_NIbU0bU7qMh2dO1NE5PbHq4Od6jF2leyRWeH8JwU1r2RNYSIe2LEjWw1qTugVYshY-LgPhffVpjU_U7Xy04ciSd2NO2J8gYoqoil7ghPvPVGsDM2eSKUQ5XCTr3vc_YBqE301-KHv1UCm5qbuXYQKi4nbntmrSP8geBTb7d8hVDFhPLFjHhklPes0vxBwvmrvJVkfC6eYSbRFd1V156EEekXnmN-yK9-fo94BmdT0124bGFVaTuBYDsmCRC5OPAYnxPwBq2qmrRwQOdiSyoLWoKVAllNktqFIL7SApPeVg'

grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "NodeScanTimeout": 500, "PortScanTimeout": 50, "MACAddress": "" }' -plaintext localhost:15123 com.pojtinger.felix.liwasc.NodeAndPortScanService.StartNodeScan
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext localhost:15123 com.pojtinger.felix.liwasc.NodeAndPortScanService.SubscribeToNodeScans
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{"ID": 11}' localhost:15123 com.pojtinger.felix.liwasc.NodeAndPortScanService.SubscribeToNodes
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{"ID": 11}' localhost:15123 com.pojtinger.felix.liwasc.NodeAndPortScanService.SubscribeToPortScans
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{"ID": 11}' localhost:15123 com.pojtinger.felix.liwasc.NodeAndPortScanService.SubscribeToPorts
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext localhost:15123 com.pojtinger.felix.liwasc.MetadataService.GetMetadataForScanner
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "MACAddress": "02:42:64:96:47:8a" }' localhost:15123 com.pojtinger.felix.liwasc.MetadataService.GetMetadataForNode
grpcurl -H "X-Liwasc-Authorization: ${AUTH}" -plaintext -d '{ "PortNumber": 7005, "TransportProtocol": "tcp" }' localhost:15123 com.pojtinger.felix.liwasc.MetadataService.GetMetadataForPort
