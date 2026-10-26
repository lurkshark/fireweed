package content

// These placeholders make the public interface compile during the evals phase.
func Version() string { return "" }

func Help() string { return "" }

func Playbook(string) (string, error) { return "", nil }

func Explain(string) (string, error) { return "", nil }
