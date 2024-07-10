export default class ArrayProblems {
	private getUpdatedEmail (email: string): string {
		const [local, domain] = email.split("@");

		let updated_local = "";

		outer:
		for (let k of local) {
			switch (k) {
			case "." :
				break;
			case "+" :
				break outer;
			default: 
				updated_local += k;
				break;
			}
		}

		const new_mail = updated_local + "@" + domain;

		

		return new_mail;
	}

	validEmails (emails: string[]): number {
		const emails_map: Record<string, boolean> = {};
		let count = 0;

		for (const email of emails) {
			const updated_email = this.getUpdatedEmail(email);
			if (!(updated_email in emails_map)) {
				emails_map[updated_email] = true;
				count += 1;
			}
		}

		

		return count;
	}
}
