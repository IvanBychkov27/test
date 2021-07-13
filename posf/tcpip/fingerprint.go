package main

func fingerprint() {

}

//perfectScore = 10
//scores = []
//for i, entry in enumerate(dbList):
//score = 0
//# @TODO: consider `ip_tll`
//# @TODO: consider `tcp_window_scaling`
//# check IP DF bit
//if entry['ip_df'] == fp['ip_df']:
//score += 1
//# check IP MF bit
//if entry['ip_mf'] == fp['ip_mf']:
//score += 1
//# check TCP window size
//if entry['tcp_window_size'] == fp['tcp_window_size']:
//score += 1.5
//# check TCP flags
//if entry['tcp_flags'] == fp['tcp_flags']:
//score += 1
//# check TCP header length
//if entry['tcp_header_length'] == fp['tcp_header_length']:
//score += 1
//# check TCP MSS
//if entry['tcp_mss'] == fp['tcp_mss']:
//score += 1.5
//# check TCP options
//if entry['tcp_options'] == fp['tcp_options']:
//score += 3
//else:
//# check order of TCP options (this is weaker than TCP options equality)
//orderEntry = ''.join([e[0] for e in entry['tcp_options'].split(',') if e])
//orderFp = ''.join([e[0] for e in fp['tcp_options'].split(',') if e])
//if orderEntry == orderFp:
//score += 2
//
//scores.append({
//'i': i,
//'score': score,
//'os': entry.get('os', {}).get('name'),
//})
