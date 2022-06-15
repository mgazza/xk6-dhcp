import dhcp from 'k6/x/dhcp'

export default function () {
    console.log("mooo")
    const client = dhcp.new("en0")

    const res =client.Exchange()
    console.log(res)
}